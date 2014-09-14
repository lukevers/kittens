// Change User Admin Settings is a function that
// changes a users administrative setting.
function ChangeUserAdminSetting()
{
	// Show enable/disable text on hover
	$('.switch_admin').on({
		mouseenter: function() {
			$(this).children().eq(0).stop(true, true).fadeIn(200);
		},
		mouseleave: function() {
			$(this).children().eq(0).stop(true, true).fadeOut(200);
		}
	});
}

// Fake Checkboxes is a function that's used instead of
// using html checkboxes. It uses a text input, and
// some fontawesome icons.
function FakeCheckboxs()
{
	$('.checkbox').bind('click', function() {
		if ($(this).hasClass('fa-times')) {
			$(this).removeClass('fa-times').addClass('fa-check');
			$('#'+$(this).data('for')).val('true');
		} else if ($(this).hasClass('fa-check')) {
			$(this).removeClass('fa-check').addClass('fa-times');
			$('#'+$(this).data('for')).val('false');
		}
	});
}
