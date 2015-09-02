var login = new Vue({
    el: '#login-form',
    data: {
        username: '',
        password: '',
        token: '',
        lock: false,
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
