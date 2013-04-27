// Cat Cache!
//
// log.js

// appendLog appends to the log by prepending the timestamp and
// sending the message to the console.
function appendLog(message) {
	// timestamp returns the current date and time in a string form,
	// which is used for logging with the console.
	function timestamp() {
		var date = new Date();
		// getCurrentDate returns the current date in a String form.
		function getCurrentDate() {
			return (String(date.getMonth())+"."+String(date.getDay())+"."+String(date.getFullYear()));
		}
		// getCurrentTime returns the current time in a String form.
		function getCurrentTime() {
			return (String(date.getHours())+":"+String(date.getMinutes())+":"+String(date.getSeconds()));
		}
		return (getCurrentDate()+" at "+getCurrentTime()+" ==> ");
	}
	console.log(timestamp()+message);
}

exports.appendLog = appendLog;
