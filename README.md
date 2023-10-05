# Golang_Practice

goroutine: process go function in background (can increase performance in correct situation) </br></br>

interface: define funtionality (dog and cat are both animals) </br> </br>

waitGroup: set a number, count down with "Done" method, block untill count to 0 (may return in main if no Wait()) </br></br>

channel: use to exchange data between goroutine (channel vs wait + slices.Append vs normal loop) (race condition) </br></br>

multiChannel: design style (pull all sender in background, include "done" sender, listen from main) </br></br>

select: case trigger by channel, randomly pick a case when cases happen simultaneously (overlap vs not overlapping)</br></br>

select2: small exercise to send/print all value before program end. </br></br>

buffer: buffer vs unbuffer, unbuffer channel will need to send out before program end. buffer will have a setted limit</br></br>

mutex: set up mutex lock to prevent racing condition </br></br>

context: use to close multiple background process </br></br>

closeChannel: close channel to signify listeners that the channel is closed </br></br>

testBench: Show an example usage of test/benchmark (syntax:go test -v or go test -v -bench=. )</br></br>

connect4bot: small connect4 bot I created using goroutine to calculate the best move. </br></br>

Keywords:</br>
break        break out loop</br>
default      default case for switch/select</br>
func         function</br>
interface    provide method signatures</br>
switch 	     another style of if/else  /check from top down</br>
select	     like switch but deals with channel/ if both condition is met, randomly pick one</br>
case         use with select and switch</br>
defer        run after function return</br>
go           run in background/goroutine</br>
map          map data structure</br>
struct       structure/ class like</br>
chan         channel</br>
else         if/else</br>
goto         goto tag</br>
package      package name</br>
const        cannot be modified after declare/ usually use to assign scalar values</br>
fallthrough  use with switch/ if one condition is met, go to the next one</br>
if           if/else</br>
range        range through slice/channel...etc</br>
type 	     create type</br>
continue     end current loop and start from next iteration</br>
for          for loop</br>
import       import package</br>
return       return value</br>
var	         use for variable declaration</br></br>

Predeclared identifiers:</br>
append	     combine slices/arrays</br>
bool	     bool variable</br>
byte		 alias for the unsigned integer 8 type ( uint8 ). Range from 0-255</br>
cap		     get array capacity</br>
close		 close a channel</br>
complex	     use for complex number</br>
complex64	 complex number float32/float32</br>
complex128	 complex number float64/float64</br>
uint16		 ranging from 0 to 65535</br>
copy		 copy one slice into another slice</br>
false		 flase</br>
float32		 float variable with 32bit </br>
float64		 float variable with 64bit </br>
imag		 get imaginary number from complex number</br>
int		     integer 32 bit</br>
int8		 -128 to 127</br>
int16		 -32,768 to 32,767</br>
uint32		 0 to 4,294,967,295</br>
int32		 -2,147,483,648 to 2,147,483,647</br>
int64		 -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807</br>
iota		 integer to ASCII</br>
len		     length of slice/array</br>
make		 create object in memory/ usually use on channel, map, slice</br>
new		     create object in memory, return pointer, automatically set default values.</br>
nil		     zero value for pointers, interfaces, maps, slices, channels and function types</br>
panic		 create a panic</br>
uint64		 0 to 18,446,744,073,709,551,615</br>
print		 will write to standard error, use fmt print instead</br>
println		 will write to standard error, use fmt print instead</br>
real		 gets the real number from complex number</br>
recover	     catch panic, should be used in a defer function.</br>
string		 string type</br>
true		 true</br>
uint		 </br>
uint8        </br>
uintptr		 integer representation of a memory address</br>
