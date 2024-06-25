# Go direct read
reads straight from /dev/input for keyboard events

## Compilation
To compile for the host system 
```go build```


to cross compile for arm  
```env GOOS=linux GOARCH=arm64 go build -o pi_main ```

## Key code information
### Make up
Information is read in 24 byte chunks 
* The first 16 bytes are for the time stamp
* The next 2 bytes are the type of event
* the next 2 bytes are the code
* The final 4 bytes are the value 

### Values
* 0 - indicates the key being released
* 1 - indicates a key being depressed 
* 2 - indicates a repeat of the key being sent 

### Types
    EV_SYN (0x00):
        Represents a synchronization event. It is typically used to synchronize multiple events.

    EV_KEY (0x01):
        Represents a keyboard key event. It is used for key press and key release events. The "code" field specifies the specific key that was pressed or released, and the "value" field indicates the event type (press/release).

    EV_REL (0x02):
        Represents a relative axis event. It is used for relative movements, such as mouse movements or scrolling. The "code" field specifies the specific axis, and the "value" field indicates the relative movement amount.

    EV_ABS (0x03):
        Represents an absolute axis event. It is used for absolute positions, such as touchscreen or joystick positions. The "code" field specifies the specific axis, and the "value" field indicates the absolute position.

    EV_MSC (0x04):
        Represents miscellaneous events. It is used for various device-specific or miscellaneous events that do not fit into other categories.

    EV_SW (0x05):
        Represents switch events. It is used for events related to switches or buttons on the input device, such as power buttons or lid switches.

    EV_LED (0x11):
        Represents LED events. It is used to control the state of LEDs on the input device, such as Caps Lock or Num Lock LEDs.

    EV_SND (0x12):
        Represents sound events. It is used for sound-related events, such as volume changes or playback control.

    EV_REP (0x14):
        Represents autorepeat events. It is used to indicate autorepeat behavior for certain keys.

    EV_FF (0x15):
        Represents force feedback events. It is used to control force feedback devices.