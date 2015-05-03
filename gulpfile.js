/*
|--------------------------------------------------------------------------
| Dependencies
|--------------------------------------------------------------------------
*/

var gulp        = require('gulp');
var less        = require('gulp-less');
var concat      = require('gulp-concat');
var watch       = require('gulp-watch');
var livereload  = require('gulp-livereload');
var minifyCSS   = require('gulp-minify-css');
var uglify      = require('gulp-uglify');
var runSequence = require('run-sequence');

/*
|--------------------------------------------------------------------------
| Variables
|--------------------------------------------------------------------------
*/

// Less
var _less  = 'app/assets/less/';
var less_  = 'app/assets/css/';

// CSS
var _css  = 'app/assets/css/';
var css_  = 'public/assets/css/';

// JavaScript
var _js   = 'app/assets/js/';
var js_   = 'public/assets/js/';

// Fonts
var fonts = 'public/assets/fonts/';

// Bower 
var bower = 'bower_components/';

/*
|--------------------------------------------------------------------------
| Less Files
|--------------------------------------------------------------------------
*/

var less_files = [
	// -- Add Less files from assets -- //

	'main.less',

	// -- End Less files from assets -- //
].map(function(str) { return _less + str; });

/*
|--------------------------------------------------------------------------
| CSS Files
|--------------------------------------------------------------------------
*/

// Bower CSS files
var css_bower = [  
	// -- Add JS files from bower -- //

	'font-awesome/css/font-awesome.min.css',

	// -- End JS files from bower -- //
].map(function(str) { return bower + str });

// Asset CSS files
var css_files = css_bower.concat([
	// -- Add CSS files from assets -- //

	'bootstrap.css',
	'main.css',

	// -- End CSS files from assets -- //
].map(function(str) { return _css + str }));

/*
|--------------------------------------------------------------------------
| JavaScript Files
|--------------------------------------------------------------------------
*/

// Bower JS files
var js_bower = [  
	// -- Add JS files from bower -- //

	'jquery/dist/jquery.min.js',
	'bootstrap/dist/js/bootstrap.min.js',
	'moment/min/moment.min.js',
	'livestampjs/livestamp.min.js',

	// -- End JS files from bower -- //
].map(function(str) { return bower + str });

// Asset JS files
var js_files = js_bower.concat([
	// -- Add JS files from assets -- //

	'channel.js',
	'server.js',
	'settings.js',
	'users.js',
	'main.js',

	// -- End JS files from assets -- //
].map(function(str) { return _js + str }));

/*
|--------------------------------------------------------------------------
| Bootstrap Task
|--------------------------------------------------------------------------
*/

gulp.task('bootstrap', function() {
	gulp.src([bower + 'bootstrap/less/bootstrap.less',
	          _less + 'bootstrap/override.less'])
		.pipe(less())
		.pipe(concat('bootstrap.css'))
		.pipe(gulp.dest(less_));
});

/*
|--------------------------------------------------------------------------
| Less Task
|--------------------------------------------------------------------------
*/

gulp.task('less', function() {
	return gulp.src(less_files)
		.pipe(less())
		.pipe(concat('main.css'))
		.pipe(gulp.dest(less_));
});

/*
|--------------------------------------------------------------------------
| CSS Task
|--------------------------------------------------------------------------
*/

gulp.task('css', function() {
	return gulp.src(css_files)
		.pipe(concat('styles.css'))
		.pipe(minifyCSS({keepSpecialComments: 0}))
		.pipe(gulp.dest(css_));
});

/*
|--------------------------------------------------------------------------
| JavaScript Task
|--------------------------------------------------------------------------
*/

gulp.task('js', function() {
	return gulp.src(js_files)
		.pipe(concat('scripts.js'))
		.pipe(uglify())
		.pipe(gulp.dest(js_));
});

/*
|--------------------------------------------------------------------------
| Copy Task
|--------------------------------------------------------------------------
*/

gulp.task('copy', function() {
	// Font Awesome
	gulp.src(bower + 'font-awesome/fonts/*')
		.pipe(gulp.dest(fonts));
});


/*
|--------------------------------------------------------------------------
| Watch Task
|--------------------------------------------------------------------------
*/

gulp.task('watch', function() {
	// Turn on Live Reload
	livereload.listen();

	// Watch for Less changes
	gulp.watch([_less + '*.less'], function() {
		runSequence('less', 'css');
	}).on('change', livereload.changed);

	// Watch for JavaScript Changes
	gulp.watch([_js + '*.js'], ['js']).on('change', livereload.changed);

	// Watch for our Bootstrap override file changes
	gulp.watch([_less + 'bootstrap/override.less'], ['bootstrap']).on('change', livereload.changed);

	// Watch for HTML changes
	gulp.watch(['app/views/*.html']).on('change', livereload.changed);
})

/*
|--------------------------------------------------------------------------
| Default Task
|--------------------------------------------------------------------------
*/

gulp.task('default', function() {
	runSequence('bootstrap', 'less', ['js', 'css', 'copy']);
});
