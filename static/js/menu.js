const onMenuToggleClick = () => {
	const menu = document.querySelector('.menu-inner')
	if (!menu) {
		console.error('Failed to open menu because element with menu inner class was not found')
		return
	}

	menu.classList.toggle('toggled')
}

const attachMenuHandlers = () => {
	const menuToggle = document.querySelector('.menu-toggle')
	if (!menuToggle) {
		console.warn('Element with menu toggle class not found, skipping menu handler attachments')
		return
	}

	menuToggle.addEventListener('click', onMenuToggleClick)
}

attachMenuHandlers()
