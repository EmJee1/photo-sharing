const handleApiError = (resBody) => {
	const message = resBody.error || 'Er is een onbekende fout opgetreden'
	UIkit.notification({message, status: 'danger'})
}

const dynamicallyImportCss = (url) => {
	const el = document.createElement('link')
	el.rel = 'stylesheet'
	el.href = url
	document.head.append(el)
}

let cropperJsLoaded = false
const loadCropperJs = async () => {
	if (cropperJsLoaded) {
		return console.debug('Not loading cropperJs because it has already been loaded in this session')
	}

	dynamicallyImportCss('https://cdnjs.cloudflare.com/ajax/libs/cropperjs/1.5.13/cropper.min.css')
	await import('https://cdnjs.cloudflare.com/ajax/libs/cropperjs/1.5.13/cropper.min.js')
	cropperJsLoaded = true
}

let cropper
const createCropper = async (imgElement) => {
	await loadCropperJs()
	cropper = new Cropper(imgElement, {
		aspectRatio: 1,
		zoomable: false,
	})
}

const attachCreatePostListeners = () => {
	const postImageInput = document.getElementById('post-image-input')
	const postImageContainer = document.getElementById('post-image-container')
	const postImagePreview = document.getElementById('post-image-preview')
	const postImageForm = document.getElementById('post-image-form')

	if (!postImageInput || !postImageContainer || !postImagePreview || !postImageForm) {
		return console.debug('No create-post-modal found on this page, skipping event listener attachment')
	}

	postImageInput.addEventListener('input', async e => {
		const [file] = e.target.files
		if (!file) {
			return console.debug('No cropper instance instantiated because no file was selected')
		}

		const previewUrl = URL.createObjectURL(file)
		postImageContainer.style.display = 'none'
		postImagePreview.style.display = 'block'

		const previewImage = postImagePreview.querySelector('img')
		previewImage.src = previewUrl
		await createCropper(previewImage)
	})

	postImageForm.addEventListener('submit', async e => {
		e.preventDefault()
		cropper.getCroppedCanvas().toBlob(async blob => {
			const formData = new FormData(e.target)
			formData.set('image', blob, 'img.png')

			try {
				await fetch('/post', {
					method: 'POST',
					body: formData,
				})
			} catch (e) {
				console.error(e)
			}
		})
	})
}
attachCreatePostListeners()

const setTabOnPageLoad = () => {
	const tabInUrl = new URL(window.location).searchParams.get('tab')
	const tabElement = document.querySelector('[data-uk-tab]')
	if (!tabElement) {
		return console.debug('Not setting tab because no tab element present in DOM')
	}

	UIkit.tab(tabElement).show(tabInUrl ?? 0)
}
setTabOnPageLoad()

UIkit.util.on('[data-url-tab]', 'show', (e) => {
	const tabIndex = e.target.getAttribute('data-tab-index')
	const url = new URL(window.location)
	url.searchParams.set('tab', tabIndex)
	history.replaceState({}, '', url)
})


const attachInviteFormListeners = () => {
	const inviteForm = document.querySelector('[data-form="invite"]')
	if (!inviteForm) {
		return console.debug('No invite-form found on this page, skipping event listener attachment')
	}

	inviteForm.addEventListener('submit', async e => {
		e.preventDefault()
		const resp = await fetch('/invite', {
			method: 'POST',
			body: new FormData(e.target)
		})
		const body = await resp.json()
		if (!body.ok) {
			handleApiError(body)
		}
	})
}
attachInviteFormListeners()
