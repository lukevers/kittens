// Fix Channel Data is a function that is called ASAP 
// because it needs to figure out what information goes 
// where, and what channel we're looking at. Since HTTP 
// Fragments (#hashes at the end of URLs) are not sent to 
// the server we can't parse what channel we're looking at 
// from the server. In order to fix that we look for a 
// hidden `div` with the id name of `location.hash` in it.
// It will contain data-attributes for every part of the 
// Channel struct that we need here on the web interface.
function FixChannelData()
{
	// Get our correct information
	var channel = location.hash;

	// Replace `#` with `\\#`
	channel = channel.replace(/#/g, '\\#');

	// Update the name of the channel
	$('#name').text($('#'+channel).data('name'));
}
