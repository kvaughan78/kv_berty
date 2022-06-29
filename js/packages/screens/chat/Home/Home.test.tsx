import { mockServices } from '@berty/utils/testing/mockServices.test'
import { renderScreen } from '@berty/utils/testing/renderScreen.test'

import { Home } from './Home'

test('Chat.Home renders correctly', async () => {
	await mockServices()

	const { toJSON } = renderScreen('Chat.Home', Home)
	expect(toJSON()).toMatchSnapshot()
})
