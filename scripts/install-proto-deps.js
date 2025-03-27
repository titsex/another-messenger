#!/usr/bin/env node
const { execSync } = require('node:child_process')

const platform = process.platform

function isGoInstalled() {
	try {
		execSync('go version', { stdio: 'ignore' })
		return true
	} catch {
		return false
	}
}

async function installGoWindows() {
	if (platform !== 'win32') return

	console.log('Checking Go installation...')
	if (isGoInstalled()) {
		console.log('Go is already installed')
		return
	}

	console.log('Installing Go for Windows...')
	try {
		execSync(
			'curl -o goinstaller.msi https://dl.google.com/go/go1.24.1.windows-amd64.msi',
			{ stdio: 'inherit' }
		)

		execSync('msiexec /i goinstaller.msi /quiet', { stdio: 'inherit' })

		await new Promise((resolve) => setTimeout(resolve, 10000))
		execSync('setx PATH "%PATH%;C:\\Program Files\\Go\\bin"', { stdio: 'inherit' })

		console.log('Go installed successfully')
	} catch (error) {
		console.error('Failed to install Go:', error)
		process.exit(1)
	}
}

async function main() {
	try {
		if (platform === 'win32') {
			await installGoWindows()
			console.log('Windows detected, using PowerShell...')
			execSync('pnpm install:proto-deps:win', { stdio: 'inherit' })
		} else {
			console.log('Unix-like system detected, using bash...')
			execSync('pnpm install:proto-deps', { stdio: 'inherit' })
		}
	} catch (error) {
		console.error('Installation failed:', error)
		process.exit(1)
	}
}

main()
