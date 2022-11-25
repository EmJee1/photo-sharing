import {css, html, LitElement} from 'https://cdn.jsdelivr.net/gh/lit/dist@2/core/lit-core.min.js'

export class SimpleGreeting extends LitElement {
	static get styles() {
		return css`p { color: blue }`;
	}

	static get properties() {
		return {
			name: {
				type: String
			}
		}
	}

	constructor() {
		super();
	}

	render() {
		return html`<p>Hello, ${this.name}!</p>`;
	}
}

customElements.define('simple-greeting', SimpleGreeting);
