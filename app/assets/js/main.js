var loc = location.href.split('/')[3];

$(function() {
	switch (loc) 
	{
		case 'server':
			JoinChannel();
			PartChannel();
			EnableDisable();
			break;
		default:
			break;
	}
});

