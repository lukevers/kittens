var settings = new Vue({
    el: '#settings-page',
    data: {
        username:     '',
        password:     '',
        email:        '',
        qr:           '',
        token:        '',
        lock:         false,
        errors:       [],
        modalErrors:  [],
        disable:      false,
        disabletext:  'Disable 2FA',
    },
    methods: {
        update: function(e) {
            e.preventDefault();

            // Lock so we don't send too many requests
            if (this.lock) return;
            this.lock = true;

            $.ajax({
                type: 'POST',
                url:  '/settings',
                data: {
                    username: this.username,
                    password: this.password,
                    email:    this.email,
                },
            }).always(function(res) {
                settings.lock = false;
                settings.$set('errors', []);
                if (res.status !== 200) {
                    var error = res.responseJSON.errors;
                    if (!ansuz.isArray(error)) {
                        // It's possible that we are only sent one error,
                        // and if so then let's convert it to an array so
                        // we can use the same function to parse the err.
                        error = [error];
                    }

                    error.map(function(err) {
                        settings.errors.push(err);
                    });
                }
            });
        },
        modal2fa: function(e) {
            e.preventDefault();
            $.ajax({
                type: 'GET',
                url:  '/settings/generate2fa',
            }).always(function(res) {
                settings.$set('modalErrors', []);
                if (res.status !== 200) {
                    var error = res.responseJSON.errors;
                    if (!ansuz.isArray(error)) {
                        error = [error];
                    }

                    error.map(function(err) {
                        settings.modalErrors.push(err);
                    });
                }

                settings.$set('qr', 'data:image/png;base64, ' + res.data);
                $('.modal').addClass('active');
            });
        },
        closeModal: function(e) {
            e.preventDefault();
            $('.modal').removeClass('active');
        },
        enable2fa: function(e) {
            e.preventDefault();
            $.ajax({
                type: 'POST',
                url:  '/settings/verify2fa',
                data: {
                    token: this.token,
                },
            }).always(function(res) {
                settings.$set('modalErrors', []);
                if (res.status !== 200) {
                    var error = res.responseJSON.errors;
                    if (!ansuz.isArray(error)) {
                        error = [error];
                    }

                    error.map(function(err) {
                        settings.modalErrors.push(err);
                    });
                } else {
                    alert('SUCCESS');
                }
            });
        },
        disable2fa: function(e) {
            e.preventDefault();
            if (!this.disable) {
                this.$set('disabletext', 'Are you sure?');
                this.disable = true;
            } else {
                $.ajax({
                    type: 'POST',
                    url:  '/settings/disable2fa',
                }).always(function(res) {
                    console.log(res);
                    if (res.status === 200) {
                        // Refresh page.
                        window.location.href = "/settings";
                    }
                });
            }
        },
    }
});
