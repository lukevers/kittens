// EnableTwoFa is a function that is called when a user clicks the
// "Enable 2FA" button on the settings page. This function enables
// (or tries to) two factor authentication on the current users
// account.
function EnableTwoFa()
{
	$('#twofa_enable').bind('click', function() {
		$.ajax({
			url: '/settings/2fa/generate',
			type: 'GET',
			success: function(data) {
				$('#qrcode').attr('src', 'data:image/png;base64, '+data);
				$('.lightbox').fadeIn(500);
			}
		});
	});
}

// DisableTwoFa is a function that is called when a user
// clicks the "Disable 2FA" button on the settings page.
// This button disables 2FA for the users account.
function DisableTwoFa()
{
	$('#twofa_disable').bind('click', function() {
		if ($(this).text() != 'Are you sure?') {
			$(this).text('Are you sure?');
			setTimeout(function() {
				$('#twofa_disable').text('Disable 2FA');
			}, 2000);
		} else {
			$.ajax({
				url: '/settings/2fa/disable',
				type: 'POST',
				success: function(data) {
					location.reload();
				}
			});
		}
	});
}

// VerifyTwoFa runs when a user clicks "verify" when
// verifying their 2FA client with a token.
function VerifyTwoFa()
{
	$('#twofa_verify').bind('click', function(e) {
		e.preventDefault();
		$.ajax({
			url: '/settings/2fa/verify',
			type: 'POST',
			data: { token: $('#token').val() },
			beforeSend: function(xhr) {
				$('#twofa_verify').html('Verifying <i class="fa fa-spin fa-refresh"></i>');
			},
			error: function(data) {
				$('#twofa_verify').html('Error <i class="fa fa-times"></i>');
				setTimeout(function() {
					$('#twofa_verify').html('Verify');
				}, 2000);
			},
			success: function(data) {
				$('#twofa_verify').html('Verified <i class="fa fa-check"></i>');
				setTimeout(function() {
					$('.lightbox').fadeOut(500, function() {
						location.reload();
					});
				}, 500);
			}
		});
	});
}

// CancelTwoFa runs when a user clicks "cancel" when
// deciding to use 2fa or not.
function CancelTwoFa()
{
	$('#twofa_cancel').bind('click', function(e) {
		e.preventDefault();
		$('.lightbox').fadeOut(500);
		$('#twofa_verify').html('Verify');
	});
}
