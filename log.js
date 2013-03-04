// Cat Cache!
//
// log.js

// The function appendLog makes appending
// The log easy without having to add the
// Timestamp function before each message
// Sent to the console.
function appendLog(message) {
	// The function timestamp returns
	// The current date and time in a
	// String form, which is used for
	// logging with the console.
	function timestamp() {
		var date = new Date();
		// The function getCurrentDate returns
		// The current date in a String form.
		function getCurrentDate() {
			return (String(date.getMonth())+"."+String(date.getDay())+"."+String(date.getFullYear()));
		}
		// The function getCurrentTime returns
		// The current time in a String form.
		function getCurrentTime() {
			return (String(date.getHours())+":"+String(date.getMinutes())+":"+String(date.getSeconds()));
		}
		return (getCurrentDate()+" at "+getCurrentTime()+" ==> ");
	}
	console.log(timestamp()+message);
}

exports.appendLog = appendLog;