package using_strings

import (
	"fmt"
	"strings"
)

func add_to(from *strings.Builder, newValue string) string {
	/*
		i was originally doing `var s strings.Builder = from`, but this is bas
		because internally, string builders have a state that tracks what's in the buffer.
		so attempting to copy a string builder will also mean you are copying the internal state
		of the stirng builder you're copying from. This can lead to unwanted behavior since
		you most likely would want a new string builder with an empty buffer/state

		The approach to solving this was to send a pointer to the function instead, that way
		we're working with the same string builder and maintain one state that is known.

		Be careful for struct methods that accept pointer receivers and not value receivers.
		the function signature for write value is: (*string.Builder).WriteString(s string) (int, error)
		this indicates that we should be using a pointer when we pass it around (maybe?? idk)
	*/
	from.WriteString(newValue)
	return from.String()
}

func Run() {
	// use the string builder when u want to keep adding to a string
	var nameBuilder strings.Builder
	nameBuilder.WriteString("Nnaemeka ")
	add_to(&nameBuilder, "Onyeokoro")
	fmt.Println(nameBuilder.String())
}
