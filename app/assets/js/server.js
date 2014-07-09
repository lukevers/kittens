// Join Channel runs when the return key is hit on the server
// page to add a new channel for the bot. Sends a POST request
// to "/server/{id}/channel/join" with the channel to join.
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
				// Send a POST request
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
// "/server/{id}/channel/part" with the channel to part.
function PartChannel()
{
	$('.channel > i').bind('click', function() {
		$(this).parent().fadeOut(500, function() {
			c = $(this)[0].children[1].innerHTML;
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

// Enable/Disable is used to show an animation on hover of the 
// "Enabled/Disabled" element. It also sends a POST request to
// "/server/{id}/enable" with a true/false variable.
function EnableDisable()
{
	// Enabled
	$('#enabled .stat-icon').bind('mouseover', function() {
		$(this).removeClass('green').addClass('red');
		$(this).children().eq(0).removeClass('fa-check').addClass('fa-times');
	});

	$('#enabled .stat-icon').bind('mouseout', function() {
		$(this).removeClass('red').addClass('green');
		$(this).children().eq(0).removeClass('fa-times').addClass('fa-check');
	});

	$('#enabled .stat-icon').bind('click', function() {
		$.ajax({
			type: 'POST',
			url: '/server/'+location.href.split('/')[4]+'/enable',
			data: {enabled: false}
		}).done(function() {
			location.reload(true);
		});
	});

	// Disabled
	$('#disabled .stat-icon').bind('mouseover', function() {
		$(this).removeClass('red').addClass('green');
		$(this).children().eq(0).removeClass('fa-times').addClass('fa-check');
	});

	$('#disabled .stat-icon').bind('mouseout', function() {
		$(this).removeClass('green').addClass('red');
		$(this).children().eq(0).removeClass('fa-check').addClass('fa-times');
	});

	$('#disabled .stat-icon').bind('click', function() {
		$.ajax({
			type: 'POST',
			url: '/server/'+location.href.split('/')[4]+'/enable',
			data: {enabled: true}
		}).done(function() {
			location.reload(true);
		});
	});
}

// View Channel Options is called when the text next to the 'x'
// for a channel is clicked. It takes you to view the channel
// options at "/server/{id}/channel/{channel}".
function ViewChannelOptions()
{
	$('.channel > .chan').bind('click', function() {
		window.location += '/channel/' + $(this)[0].firstChild.data;
	});
}
