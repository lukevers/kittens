var loc = location.href.split('/')[3];

$(function() {
	switch (loc) 
	{
		case 'server':
			switch (location.href.split('/')[5])
			{
				case 'channel':
					RedirectIfFragmented();
					break;
				default:
					break;
			}
			switch (location.href.split('/')[4])
			{
				case 'new':
					FakeCheckboxs();
					break;
				default:
					JoinChannel();
					PartChannel();
					EnableDisable();
					ViewChannelOptions();
					FakeCheckboxs();
					break;
			}
			break;
		case 'settings':
			EnableTwoFa();
			DisableTwoFa();
			VerifyTwoFa();
			CancelTwoFa();
			break;
		case 'users':
			ChangeUserAdminSetting();
			FakeCheckboxs();
			break;
		default:
			break;
	}
});

