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

var css_files = [
	// -- Add CSS files from assets -- //

	'bootstrap.min.css',
	'font-awesome.min.css',
	'main.css',

	// -- End CSS files from assets -- //
].map(function(str) { return _css + str });

/*
|--------------------------------------------------------------------------
| JavaScript Files
|--------------------------------------------------------------------------
*/

// Public JS files
var js_files = [
	// -- Add JS files from assets -- //

	'jquery.min.js',
	'bootstrap.min.js',
	'moment.min.js',
	'livestamp.min.js',
	'channel.js',
	'server.js',
	'settings.js',
	'main.js',

	// -- End JS files from assets -- //
].map(function(str) { return _js + str });

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

	// Watch for HTML changes
	gulp.watch(['app/views/*.html']).on('change', livereload.changed);
})

/*
|--------------------------------------------------------------------------
| Default Task
|--------------------------------------------------------------------------
*/

gulp.task('default', function() {
	runSequence('less', ['js', 'css']);
});
