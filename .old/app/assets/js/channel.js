// Redirect If Fragmented is a function that's called when
// a user visits "/server/{id}/channel/#{channel}" instead
// of "/server/{id}/channel/{channel}" and takes the URL
// fragment from the URL. It then URL encodes the URL and
// redirects us to the correct URL.
function RedirectIfFragmented()
{
	// Check if our URL looks like "/server/{id}/channel/".
	// We also need to make sure that the last part of the
	// array we split equals "".
	if (window.location.pathname.split('/').length == 5
		&& window.location.pathname.split('/')[4] == '') {
		// check if we have a hash. If we don't have a hash
		// it means we're just at "/server/{id}/channel/",
		// so we'll redirect back to "/server/{id}" instead.
		if (window.location.hash == '') {
			parts = window.location.pathname.split('/');
			window.location = '/server/' + parts[2];
		} else {
			path = window.location.pathname + encodeURIComponent(window.location.hash);
			window.location = path;
		}
	}
}
