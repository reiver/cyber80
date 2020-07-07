(function(){
	console.log("magic256: Hello world!")
})();

(function(){

	window.requestAnimationFrame = window.requestAnimationFrame
	                            || window.mozRequestAnimationFrame
	                            || window.webkitRequestAnimationFrame
	                            || window.msRequestAnimationFrame
	                            || function(f){return setTimeout(f, 1000/60)}


	window.cancelAnimationFrame = window.cancelAnimationFrame
	                           || window.mozCancelAnimationFrame
	                           || function(requestID){clearTimeout(requestID)}
})();

(function(){

	const canvasid = "v";

	const canvas = document.getElementById(canvasid);
	if (!canvas) {
		console.log("could not find <canvas> with id:", canvasid);
		return;
	}

	const width  = 256;
	const height = 256;

	canvas.setAttribute("width", width);
	canvas.setAttribute("height", height);
	document.title = document.title+": "+width+"×"+height;

	const ctx = canvas.getContext("2d");
	if (!ctx) {
		console.log("could get context from <canvas> with id:", canvasid);
		return;
	}

	const imageData = ctx.getImageData(0, 0, width, height);
	const buf = new ArrayBuffer(imageData.data.length);
	const buf8 = new Uint8ClampedArray(buf);
	const data = new Uint32Array(buf);


	function drawNoise(buffer) {

		for (var y = 0; y < height; ++y) {
			for (var x = 0; x < width; ++x) {

				const red   = (Math.random() * 93) & 0xff;
				const green = 0 & 0xff;
				//const green = (Math.random() * 256) & 0xff;
				const blue  = (Math.random() * 93) & 0xff;
				const alpha = 255 | 0xff;

				const rgba =
					(alpha << 24) |  // alpha
					(blue  << 16) |  // blue
					(green <<  8) |  // green
					 red;            // red

				buffer[y * width + x] = rgba;

			}
		}
	}

	function draw(timestamp){

		var timestamp = timestamp || new Date().getTime();

		// MAGIC FUNCTION: window._next
		const fn = window._next || drawNoise;

		fn(data);

		imageData.data.set(buf8);
		ctx.putImageData(imageData,0,0);

		requestAnimationFrame(draw);
	}
	requestAnimationFrame(draw);
})();

(function(){
	console.log("magic256: I am alive!")
})();
