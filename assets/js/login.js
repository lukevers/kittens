var login = new Vue({
    el: '#login-form',
    data: {
        username: '',
        password: '',
        token:    '',
        lock:     false,
        errors:   [],
    },
    methods: {
        login: function(e) {
            e.preventDefault();

            // Lock so we don't send too many requests
            if (this.lock) return;
            this.lock = true;

            $.ajax({
                type: 'POST',
                url: '/login',
                data: {
                    username: this.username,
                    password: this.password
                },
            }).always(function(res) {
                if (res.status === 200) {
                    window.location.href = (res.twofa ? '/login/2fa' : '/');
                }

                login.lock = false;
                login.$set('errors', []);
                if (res.status !== 200) {
                    var error = res.responseJSON;
                    if (!ansuz.isArray(error)) {
                        // We probably are only ever going to recieve one error here,
                        // but just in case we change this in the future, let's take
                        // an extra step now and convert this into an array if it's
                        // not already.
                        error = [error];
                    }

                    error.map(function(err) {
                        login.errors.push(err);
                    });
                }
            });
        },

        login2fa: function(e) {
            e.preventDefault();

            // Lock so we don't send too many requests
            if (this.lock) return;
            this.lock = true;

            $.ajax({
                type: 'POST',
                url: '/login/2fa',
                data: { 
                    token: this.token 
                },
            }).always(function(res) {
                if (res.status === 200) {
                    window.location.href = '/';
                }

                login.lock = false;
            });
        },
    }
});
