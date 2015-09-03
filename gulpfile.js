var gulp   = require('gulp');
var less   = require('gulp-less');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');

var LessPluginCleanCSS = require('less-plugin-clean-css'),
    LessPluginAutoPrefix = require('less-plugin-autoprefix'),
    cleancss = new LessPluginCleanCSS({ advanced: true }),
    autoprefix= new LessPluginAutoPrefix({ browsers: ["last 2 versions"] });

gulp.task('less', function() {
    return gulp.src('assets/less/main.less')
        .pipe(less({
            plugin: [autoprefix, cleancss],
        }))
        .pipe(concat('styles.css'))
        .pipe(gulp.dest('public/assets/css/'));
});

gulp.task('js', function() {
    return gulp.src([
            'assets/js/main.js',
            'assets/js/login.js',
        ])
        .pipe(concat('scripts.js'))
        .pipe(uglify())
        .pipe(gulp.dest('public/assets/js/'));
});

gulp.task('jslibs', function() {
    return gulp.src([
            'bower_components/jquery/dist/jquery.min.js',
            'bower_components/bootstrap/dist/js/bootstrap.min.js',
            'bower_components/vue/dist/vue.min.js',
        ])
        .pipe(concat('libs.js'))
        .pipe(uglify())
        .pipe(gulp.dest('public/assets/js/'));
});

gulp.task('copy', function() {
    return gulp.src([
        'bower_components/font-awesome/fonts/*',
    ])
    .pipe(gulp.dest('public/assets/fonts/'));
});

gulp.task('watch', function() {
    gulp.watch('assets/js/*.js',     ['js']);
    gulp.watch('assets/less/**/*.less', ['less']);
});

gulp.task('default', ['less', 'js', 'jslibs', 'copy']);
