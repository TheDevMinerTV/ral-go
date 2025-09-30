import { spawnSync } from 'node:child_process';
import { createWriteStream } from 'node:fs';
import { setTimeout } from 'node:timers/promises';
import { RAL } from 'ral-colors/index.js';

type RegistryNames = keyof typeof RAL;
type NormalRegistryNames = Exclude<RegistryNames, 'effect'>;

async function generate<RegistryName extends NormalRegistryNames>(registryName: RegistryName) {
	const registry = RAL[registryName];

	console.log(`generating for registry ${registryName}`);

	const fd = createWriteStream(`${registryName}.go`);
	fd.write('package ral\n\n');
	fd.write('import "image/color"\n\n');

	const variableName = registryName.slice(0, 1).toUpperCase() + registryName.slice(1);

	fd.write(`var ${variableName} = map[string]Color{\n`);
	for (const [colorID, color] of Object.entries(registry)) {
		console.log(`  generating color ${colorID}`);
		fd.write(
			`\t${JSON.stringify(colorID)}: {RGBA: color.RGBA{R: ${color.rgb.r}, B: ${color.rgb.b}, G: ${color.rgb.g}, A: 255}, Name: ${JSON.stringify(
				color.description
			)}},\n`
		);
	}
	fd.write(`}\n`);

	fd.close();

	await setTimeout(1 * 1000);

	spawnSync('go', ['fmt', `./${registryName}.go`]);
}

generate('classic');
generate('design_system');
