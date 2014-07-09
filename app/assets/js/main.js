var loc = location.href.split('/')[3];

$(function() {
	switch (loc) 
	{
		case 'server':
			switch (location.href.split('/')[5])
			{
				case 'channel':
					FixChannelData();
					break;
				default:
					JoinChannel();
					PartChannel();
					EnableDisable();
					ViewChannelOptions();
					break;
			}
			break;
		default:
			break;
	}
});

