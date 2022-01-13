const go = new Go();

WebAssembly.instantiateStreaming(fetch('fetch.wasm'), go.importObject).then(
	(result) => {
		go.run(result.instance);
	}
);

async function MyFunc(endpoint, query) {
	try {
		const response = await MyGoFunc(endpoint, query);

		// Response is in XML format
		const str = await response.text();
		const data = await new window.DOMParser().parseFromString(str, 'text/xml');
		console.log(data);

		const results = data.getElementsByTagName('uri');
		console.log(results);
		for (result of results) {
			console.log(result.textContent);
		}
	} catch (err) {
		console.error('Caught exception', err);
	}
}
