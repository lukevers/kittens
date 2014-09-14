//
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
