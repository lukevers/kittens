var loc = location.href.split('/')[3];

$(function() {
	switch (loc) 
	{
		case 'server':
			JoinChannel();
			PartChannel();
			EnableDisable();
			ViewChannelOptions();
			break;
		default:
			break;
	}
});

