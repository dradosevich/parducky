// Danny Radosevich
// UWYO CEDAR
// Directly read keyboard input from /dev/input

// https://janczer.github.io/work-with-dev-input/
// https://www.kernel.org/doc/html/v4.18/input/event-codes.html

package main

//imports
import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const NULL_CHR = byte(0)

func writeToHidg(modifier string, keypress string) {
	//sanity check the keypress
	//fmt.Println("KP is", keypress)

	//make our byte slice
	bytes_to_send := make([]byte, 8)

	//set all the bytes to be the NULL_CHR
	for i := range bytes_to_send {
		bytes_to_send[i] = NULL_CHR
	}
	// open /dev/hidg0 to write out as a keyboard
	file, err := os.OpenFile("/dev/hidg0", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Failed to open HID gadget device:", err)
		return
	}
	defer file.Close()
	// defer the closing

	//now we need to byte-ify our hex modifier
	//done using the encoding/hex library
	modifierBytes, err := hex.DecodeString(modifier)
	if err != nil {
		fmt.Println("Invalid modifier hex string:", err)
		return
	}
	//put the modifier into the first position in what we are writing out
	bytes_to_send[0] = modifierBytes[0]

	//reepeat with our keypress
	kpBytes, err := hex.DecodeString(keypress)
	if err != nil {
		fmt.Println("Invalid kp hex string:", err)
		return
	}
	bytes_to_send[2] = kpBytes[0]

	//and now we write out the byte string and release they keyboard
	report := bytes_to_send
	file.Write(report)
	//file.Write([]byte{NULL_CHR, NULL_CHR, NULL_CHR, NULL_CHR, NULL_CHR, NULL_CHR, NULL_CHR, NULL_CHR})
}

// write log
// helper function to write to the log file
// debated logging as we went along and writing in line in main, but wanted to be able to close it each time
// closing each time toa void file corruption
// does add some overhead comp wise
func writeToLog(eventType string, cod uint16, log_name string) {

	log_file, err := os.OpenFile(log_name, os.O_APPEND|os.O_WRONLY, 0666) //open a  file to write to

	if err != nil {
		panic(err)
	}
	log_file.Write([]byte(getEngFromMap(cod) + " " + eventType + " " + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "\n"))
	err = log_file.Close()
	if err != nil {
		log.Fatal(err)
	}

}

// helper function to see if element in a slice
// I wrote this code in go 1.20.6 in July 2023
// Go 1.21 set to release in August 2023 t in method
func inSlice(mod_codes []uint16, val uint16) bool {
	found := false
	for _, value := range mod_codes {
		if val == value {
			found = true
			break
		}
	}
	return found
}

func changeModifers(mod_slice *byte, pos int, val uint8) {
	//Adjust our modfier bits
	//pos should be 0<=pos<=7
	//the modifiers' bits are:
	/*
		Modifier keys utilize a unique bit to correspond with a key
		|Bit           | Key         |
		|--------------|-------------|
		|0             | left ctrl   |
		|1             | left shift  |
		|2             | left alt    |
		|3             | left GUI    |
		|4             | right ctr   |
		|5             | right shift |
		|6             | right alt   |
		|7             | right GUI   |
	*/
	//make sure we are within bounds for our byte
	if pos >= 0 && pos <= 7 {
		if val == 0 {
			// &^= bitwise AND NOT assignment
			(*mod_slice) &^= 1 << pos
		} else if val == 1 {
			(*mod_slice) |= 1 << pos
		}
	}
}

// this function is to determine the correct eventX to read from
// depending on the keyboard it may differ
// Assumes "Keyboard" in the name
func findCorrectEventX() string {
	//where the eventX are stored
	input_path := "/dev/input/"
	ev_files, err := os.ReadDir(input_path) //read all files in that directory
	if err != nil {
		log.Fatal(err)
	}
	poss_evs := []string{}       //create a slice for us to store the current eventX in
	for _, e := range ev_files { //go through all the files
		if strings.Contains(e.Name(), "event") { //only want the ones containing `event`
			fmt.Println(e.Name())
			poss_evs = append(poss_evs, e.Name()) //add them to the above slice
		}
	}
	//now we are going to use the eventX devices found to check the input names stored
	//first get our string to format
	path_check := "/sys/class/input/%s/device/name"

	//now go through
	for _, poss_ev := range poss_evs {
		out, err := exec.Command("cat", fmt.Sprintf(path_check, poss_ev)).Output()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
		if strings.Contains(strings.ToLower(string(out)), "keyboard") {
			fmt.Printf("Found a possible eventX %s\n", poss_ev)
			//found the eventX that is for a keyboard
			return poss_ev
		}
	}
	return "null"
}

func main() {
	//generate up a UUID for a file to output
	id := uuid.New()                                                                            //generate up a random UUID for the file
	date := fmt.Sprint(time.Now().Date())                                                       //get the current date, store in string
	t := fmt.Sprint(time.Now().Clock())                                                         //get current time, store in string
	date = strings.ReplaceAll(date, " ", "-")                                                   //remove spaces and replace with dashes
	t = strings.ReplaceAll(t, " ", "-")                                                         //remove spaces and replace with dashes
	log_name := fmt.Sprint(date) + "_" + t + "-" + fmt.Sprint() + "_" + fmt.Sprint(id) + ".txt" //generate up the name for our log file of presses
	log_file, err := os.OpenFile(log_name, os.O_CREATE, 0666)                                   //open a  file to write to,r create if needed

	if err != nil {
		panic(err)
	}
	err = log_file.Close()
	if err != nil {
		log.Fatal(err)
	}

	//a slice for your modifier keys so be know what to send
	mod_keys := byte(0x00)
	//define a slice of the modifer keycodes
	//left ctrl 29
	//left shift 42
	//left alt 56
	//left gui 125
	//right ctrl 97
	//right shift 54
	//right alt 100
	//right gui126
	mod_codes := []uint16{42, 54, 29, 97, 56, 100, 125, 126}

	//first we need to open the file we are reading from
	//get tthe assumed correct eventX
	evX := findCorrectEventX()
	if evX == "null" {
		panic("couldn't find the eventX")
	}
	fmt.Println(evX)
	f, err := os.Open("/dev/input/" + evX)

	//now checking to see if there is an error
	if err != nil {
		panic(err)
	} else {
		fmt.Println("opened the input stream")
	}
	//now defer the closing of the file
	defer f.Close()

	// The data  holds 24 bytes
	// the first 16 bytes are for the time stamp
	// The next 2 bytes are the type of event
	// The next 2 bytes are the code
	// The next 4 bytes are the value

	//make a slice to store our bytes in
	all_bytes := make([]byte, 24)
	//stores if shift is currently being held or not
	//shift_held := false
	//infinite loop to read the data from the file at each event
	fmt.Println("Entering loop")
	for {
		fmt.Println("reading bytes...")
		//read from the file into the slice
		f.Read(all_bytes)
		fmt.Println("read bytes")
		//get the seconds out of the slice, and decode
		//sec := binary.LittleEndian.Uint64(all_bytes[0:8])
		//usec := binary.LittleEndian.Uint64(all_bytes[8:16])
		//decode the seconds so we cans ee the time of the event
		//t := time.Unix(int64(sec), int64(usec))
		//fmt.Println(t)
		var value int32
		//get the type and the code for the event
		typ := binary.LittleEndian.Uint16(all_bytes[16:18])
		cod := binary.LittleEndian.Uint16(all_bytes[18:20])
		binary.Read(bytes.NewReader(all_bytes[20:]), binary.LittleEndian, &value)
		fmt.Printf("type: %x\ncode: %d\nvalue: %d\n", typ, cod, value)

		//first we are going to check the type of the event
		//we only care about keyboard events which == 1
		if typ == 1 {
			fmt.Printf("Our keycode was %d which relates to %s", cod, getHexFromMap(cod))

			//now we need to do two different things for when it is a press or release
			if value == 1 { //a key is being pressed
				//log_file.Write(getEngFromMap(cod) + " KeyDown " + strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
				writeToLog("KeyDown", cod, log_name)

				if inSlice(mod_codes, cod) { //one of the mod keys pressed
					//left ctrl 29		0
					//left shift 42		1
					//left alt 56		2
					//left gui 125		3
					//right ctrl 97		4
					//right shift 54	5
					//right alt 100		6
					//right gui 126		7
					//check which modifer key is pressed, and change the corresponding bit
					if cod == 29 {
						changeModifers(&mod_keys, 0, 1)
					} else if cod == 42 {
						changeModifers(&mod_keys, 1, 1)
					} else if cod == 56 {
						changeModifers(&mod_keys, 2, 1)
					} else if cod == 125 {
						changeModifers(&mod_keys, 3, 1)
					} else if cod == 97 {
						changeModifers(&mod_keys, 4, 1)
					} else if cod == 54 {
						changeModifers(&mod_keys, 5, 1)
					} else if cod == 100 {
						changeModifers(&mod_keys, 6, 1)
					} else if cod == 126 {
						changeModifers(&mod_keys, 7, 1)
					}

					//shift_held = true //don't need this anymore I think?

				} else {

					writeToHidg(fmt.Sprintf("%02X", mod_keys), getHexFromMap(cod))

				}

			} else if value == 0 { //key is released
				writeToLog("KeyUp", cod, log_name)
				if inSlice(mod_codes, cod) { //one of the shift keys being released
					//shift_held = false
					//left ctrl 29		0
					//left shift 42		1
					//left alt 56		2
					//left gui 125		3
					//right ctrl 97		4
					//right shift 54	5
					//right alt 100		6
					//right gui 126		7
					if cod == 29 {
						changeModifers(&mod_keys, 0, 0)
					} else if cod == 42 {
						changeModifers(&mod_keys, 1, 0)
					} else if cod == 56 {
						changeModifers(&mod_keys, 2, 0)
					} else if cod == 125 {
						changeModifers(&mod_keys, 3, 0)
					} else if cod == 97 {
						changeModifers(&mod_keys, 4, 0)
					} else if cod == 54 {
						changeModifers(&mod_keys, 5, 0)
					} else if cod == 100 {
						changeModifers(&mod_keys, 6, 0)
					} else if cod == 126 {
						changeModifers(&mod_keys, 7, 0)
					}
					writeToHidg(fmt.Sprintf("%02X", mod_keys), "00")
				} else { //not shift being held
					//a normal key lifted, jsut write out with what is left
					writeToHidg(fmt.Sprintf("%02X", mod_keys), "00")

				}

			}
		}
	}
}
