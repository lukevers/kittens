var loc = location.href.split('/')[3];

$(function() {
	switch (loc) 
	{
		case 'server':
			JoinChannel();
			PartChannel();
			break;
		default:
			break;
	}
});

