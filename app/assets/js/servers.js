// Join Channel runs when the return key is hit on the server
// page to add a new channel for the bot. Sends a POST request
// to "/server/{id}/channel/join"
function JoinChannel()
{
	$('#channel').bind('keypress', function(e) {
		var code = e.keyCode || e.which;
		// Enter
		if (code == 13) {
			chans = $('#channel').val().trim().split(' ');
			$('#channel').val('');
			for (i = 0; i < chans.length; i++) {
				c = (chans[i].substring(0,1) == '#') ? chans[i] : '#' + chans[i];
				ch = '<div class="channel"><i class="fa fa-times"></i>&nbsp; '+c+'</div>\n\r';
				$('#channels').append(ch);
				$.ajax({
					type: 'POST',
					url: '/server/'+location.href.split('/')[4]+'/channel/join',
					data: {channel: c}
				});
			}
			// Run PartChannel again to bind the newly created
			PartChannel()
		}
	});
}

// Part Channel is called when the 'x' is clicked next to the
// channel name on the server page. Sends a POST request to
// "/server/{id}/channel/part"
function PartChannel()
{
	$('.channel > i').bind('click', function() {
		$(this).parent().fadeOut(500, function() {
			c = $(this)[0].childNodes[2].data.trim();
			$(this).remove();
			// Send a POST request
			$.ajax({
				type: 'POST',
				url: '/server/'+location.href.split('/')[4]+'/channel/part',
				data: {channel: c}
			});
		});
	});
}
