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

connect4bot: small connect4 bot I created using goroutine to calculate the best move. </br>
