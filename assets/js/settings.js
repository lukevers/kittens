var settings = new Vue({
    el: '#settings-page',
    data: {
        username: '',
        password: '',
        email:    '',
        lock:     false,
        errors:   [],
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
                    var error = res.responseJSON;
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
    }
});
