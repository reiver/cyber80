(function(){
	console.log("magic256: Hello world!")
})();

(function(){

	window.requestAnimationFrame = window.requestAnimationFrame
	                            || window.mozRequestAnimationFrame
	                            || window.webkitRequestAnimationFrame
	                            || window.msRequestAnimationFrame
	                            || function(f){return setTimeout(f, 1000/60)};


	window.cancelAnimationFrame = window.cancelAnimationFrame
	                           || window.mozCancelAnimationFrame
	                           || function(requestID){clearTimeout(requestID)};
})();

(function(){

	const canvases = document.getElementsByTagName("canvas")
	if (!canvases) {
		console.log("cannot find any <canvas> elements.");
		return;
	}

	const canvas = canvases[0];
	if (!canvas) {
		console.log("could not get <canvas>");
		return;
	}
	canvas.setAttribute("class", "magic256");

	const width  = 256;
	const height = 256+32;

	canvas.setAttribute("width", width);
	canvas.setAttribute("height", height);

	document.title = document.title+": "+width+"×"+height;

	const ctx = canvas.getContext("2d");
	if (!ctx) {
		console.log("could get context from <canvas>");
		return;
	}

	const imageData = ctx.getImageData(0, 0, width, height);
	const arrayBuffer = new ArrayBuffer(imageData.data.length);
	const uint8ClampedArray = new Uint8ClampedArray(arrayBuffer);

	function drawNoise(buffer) {

		const data = new Uint32Array(buffer);

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

				data[y * width + x] = rgba;

			}
		}
	}

	function draw(timestamp){

		var timestamp = timestamp || new Date().getTime();

		// MAGIC FUNCTION: window._next
		const fn = window._next || drawNoise;

		fn(arrayBuffer);

		imageData.data.set(uint8ClampedArray);
		ctx.putImageData(imageData,0,0);

		requestAnimationFrame(draw);
	}
	requestAnimationFrame(draw);
})();

(function(){
	console.log("magic256: I am alive!");
})();
