import fs from 'node:fs/promises';
const seedData = await fs.readFile('seeds.csv', 'utf-8');
