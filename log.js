// Cat Cache!
//
// log.js

// appendLog appends to the log by prepending the timestamp and
// sending the message to the console.
function appendLog(message) {
	console.log(getCurrentDate()+" at "+getCurrentTime()+" ==> "+message);
}

// getCurrentDate returns the current date in a String form.
function getCurrentDate() {
	var date = new Date();
	return (String(date.getMonth())+"."+String(date.getDay())+"."+String(date.getFullYear()));
}

// getCurrentTime returns the current time in a String form.
function getCurrentTime() {
	var date = new Date();
	return (String(date.getHours())+":"+String(date.getMinutes())+":"+String(date.getSeconds()));
}

exports.appendLog = appendLog;
