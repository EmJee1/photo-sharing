const inviteUserToGroup = async (groupId, inviteeEmail) => {
	const formData = new FormData()
	formData.append('inviteeEmail', inviteeEmail)

	await fetch(`/group/${groupId}/invite`, {
		method: 'POST',
		body: formData,
	})
}

const attachInviteFormListeners = () => {
	const inviteForm = document.querySelector('form[data-invite-form]')
	if (!inviteForm) {
		return console.debug('No invite-form found on this page, skipping event listener attachment')
	}

	inviteForm.addEventListener('submit', async (e) => {
		e.preventDefault()
		const groupId = inviteForm.getAttribute('data-group')
		const inviteeEmail = new FormData(e.target).get('email')

		try {
			await inviteUserToGroup(groupId, inviteeEmail);
			window.location.reload()
		} catch (err) {
			console.error(err);
		}
	})
}

attachInviteFormListeners()
