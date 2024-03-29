const handleApiError = (resBody) => {
	if (resBody?.ok) {
		return
	}
	const message = resBody?.error || 'Er is een onbekende fout opgetreden'
	UIkit.notification({message, status: 'danger'})
}

const apiRequest = async (url, formData, method = 'GET') => {
	let body
	try {
		const resp = await fetch(`/api/${url}`, {method, body: formData})
		body = await resp.json()
	} catch (_) {
		handleApiError()
		return
	}

	if (!body.ok) {
		handleApiError(body)
		return
	}

	window.location.reload()
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

	postImageForm.addEventListener('submit', e => {
		e.preventDefault()
		const lottiePlayer = document.querySelector('[data-post-upload-loader]')
		lottiePlayer.load('https://assets2.lottiefiles.com/temp/lf20_xYfV1x.json')
		document.querySelector('[data-post-upload-form]').style.display = 'none'
		cropper.getCroppedCanvas().toBlob(async blob => {
			const formData = new FormData(e.target)
			formData.set('image', blob, 'img.png')
			// give some time for the animation to show
			setTimeout(async () => {
				await apiRequest('post', formData, 'POST')
			}, 350)
		})
	})
}
attachCreatePostListeners()

const attachCreateCommentListener = () => {
	const forms = document.querySelectorAll('[data-comment-post]')
	forms.forEach(form => {
		form.addEventListener('submit', async (e) => {
			e.preventDefault()
			await apiRequest('comment', new FormData(e.target), 'POST')
		})
	})
}
attachCreateCommentListener()

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
	if (!tabIndex) {
		// prevent setting a tab value of 'null' if no data-tab-index attribute is set
		return
	}
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
		await apiRequest('invite', new FormData(e.target), 'POST')
	})
}
attachInviteFormListeners()

const attachPostLikeButtonListeners = () => {
	const likeButtons = document.querySelectorAll('[data-like-post]')
	likeButtons.forEach(btn => {
		btn.addEventListener('click', async () => {
			const postId = btn.getAttribute('data-like-post')
			const likeCountEl = btn.querySelector('[data-like-count]')
			const crrntLikes = Number(likeCountEl.innerText)
			const formData = new FormData()
			formData.append('postId', postId)
			const resp = await fetch('/api/like', {method: 'POST', body: formData})
			const body = await resp.json()
			handleApiError(body)
			if (body.ok) {
				if (btn.classList.contains('active')) {
					likeCountEl.innerText = crrntLikes - 1
				} else {
					likeCountEl.innerText = crrntLikes + 1
				}
				btn.classList.toggle('active')
			}
		})
	})
}
attachPostLikeButtonListeners()

const attachDeleteListeners = () => {
	const deleteCommentBtns = document.querySelectorAll('[data-delete-comment]')
	deleteCommentBtns.forEach(btn => {
		btn.addEventListener('click', async e => {
			e.preventDefault()
			const formData = new FormData()
			formData.append('commentId', btn.getAttribute('data-delete-comment'))
			await apiRequest('comment', formData, 'DELETE')
		})
	})

	const deletePostBtns = document.querySelectorAll('[data-delete-post]')
	deletePostBtns.forEach(btn => {
		btn.addEventListener('click', async e => {
			e.preventDefault()
			const formData = new FormData()
			formData.append('postId', btn.getAttribute('data-delete-post'))
			await apiRequest('post', formData, 'DELETE')
		})
	})
}
attachDeleteListeners()

const attachKickUserListeners = () => {
	const kickUserBtns = document.querySelectorAll('[data-kick-user]')
	kickUserBtns.forEach(btn => {
		btn.addEventListener('click', async () => {
			const formData = new FormData()
			formData.append('userId', btn.getAttribute('data-user'))
			formData.append('groupId', btn.getAttribute('data-group'))
			await apiRequest('kick', formData, 'POST')
		})
	})
}
attachKickUserListeners()

const attachMakeAdminListeners = () => {
	const makeAdminBtns = document.querySelectorAll('[data-make-admin]')
	makeAdminBtns.forEach(btn => {
		btn.addEventListener('click', async () => {
			const formData = new FormData()
			formData.append('userId', btn.getAttribute('data-user'))
			formData.append('groupId', btn.getAttribute('data-group'))
			await apiRequest('promote', formData, 'POST')
		})
	})
}
attachMakeAdminListeners()

const respondToInvite = async (accept, inviteId) => {
	clearCacheInvites()
	const formData = new FormData()
	formData.append('inviteId', inviteId)
	formData.append('action', accept ? 'accept' : 'reject')
	await apiRequest('invite/respond', formData, 'POST')
}

const createNotifcations = (invites) => {
	invites.forEach(invite => {
		const li = document.createElement('li')
		const p = document.createElement('p')
		p.classList.add('uk-margin-small-bottom')
		p.innerText = `Je bent uitgenodigd voor de groep: ${invite.Group.Name}`
		const actions = document.createElement('div')
		actions.classList.add('uk-flex', 'uk-flex-between')
		const accept = document.createElement('button')
		accept.classList.add('uk-button', 'uk-button-primary', 'uk-button-small')
		accept.innerText = 'Accepteren'
		accept.addEventListener('click', () => respondToInvite(true, invite.ID))
		const reject = document.createElement('button')
		reject.classList.add('uk-button', 'uk-button-danger', 'uk-button-small')
		reject.innerText = 'Afwijzen'
		reject.addEventListener('click', () => respondToInvite(false, invite.ID))
		actions.append(accept, reject)
		li.append(p, actions)
		document.querySelector('[data-notifications-list]').append(li)
	})
}

const setCacheInvites = (invites) => {
	sessionStorage.setItem('invites:cache-date', Date.now().toString())
	sessionStorage.setItem('invites', JSON.stringify(invites))
}
const clearCacheInvites = () => {
	sessionStorage.removeItem('invites:cache-date')
	sessionStorage.removeItem('invites')
}
const getCacheInvites = () => {
	const cacheDate = sessionStorage.getItem('invites:cache-date')
	const invites = sessionStorage.getItem('invites')

	// invalidate cache if the previous fetch is > 30 seconds ago
	if (cacheDate && (Date.now() - cacheDate) > 30_000) {
		return clearCacheInvites()
	}

	if (invites) {
		return JSON.parse(invites)
	}
}

const fetchInvites = async () => {
	const resp = await fetch('/api/invite')
	const body = await resp.json()
	return body.invites
}

const getInvites = async () => {
	if (['/login', '/register'].includes(window.location.pathname)) {
		clearCacheInvites()
		return
	}

	let invites = getCacheInvites()
	if (!invites) {
		invites = await fetchInvites()
		setCacheInvites(invites)
	}

	createNotifcations(invites)
}
void getInvites()
