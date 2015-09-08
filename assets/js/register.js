var register = new Vue({
    el: '#register-form',
    data: {
        username: '',
        password: '',
        email:    '',
        lock:     false,
        errors:   [],
    },
    methods: {
        register: function(e) {
            e.preventDefault();

            // Lock so we don't send too many requests
            if (this.lock) return;
            this.lock = true;

            $.ajax({
                type: 'POST',
                url: '/register',
                data: {
                    username: this.username,
                    password: this.password,
                    email:    this.email,
                },
            }).always(function(res) {
                if (res.status === 200) {
                    window.location.href = '/login';
                }

                register.lock = false;
                register.$set('errors', []);
                if (res.status !== 200) {
                    var error = res.responseJSON.errors;
                    if (!ansuz.isArray(error)) {
                        error = [error];
                    }

                    error.map(function(err) {
                        register.errors.push(err);
                    });
                }
            });
        },
    }
});
