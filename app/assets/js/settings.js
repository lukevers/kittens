// EnableTwoFa is a function that is called when a user clicks the
// "Enable 2FA" button on the settings page. This function enables
// (or tries to) two factor authentication on the current users
// account.
function EnableTwoFa()
{
	$('#twofa_enable').bind('click', function() {
		$('#twofa_enable, #message').fadeOut(500, function() {
			$.ajax({
				url: '/settings/2fa/generate',
				type: 'GET',
				success: function (data) {
					$('#qrcode').attr('src', 'data:image/png;base64, '+data);
					$('#qr').fadeIn(500);
				}
			});
		});
	});
}
