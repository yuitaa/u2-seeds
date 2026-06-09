import fs from 'node:fs/promises';

interface Mapping {
  tasks: Map<string, number>;
  extensions: Map<string, number>;
  wagons: Map<string, number>;
  engines: Map<string, number>;
  versions: Map<string, number>;
}

const mappingData = await fs.readFile('mapping.json', 'utf-8');
const mapping = JSON.parse(mappingData) as Mapping;

export const tasksReverse = new Map(
  Array.from(mapping.tasks, ([k, v]) => [v, k]),
);
export const extensionsReverse = new Map(
  Array.from(mapping.extensions, ([k, v]) => [v, k]),
);
export const wagonsReverse = new Map(
  Array.from(mapping.wagons, ([k, v]) => [v, k]),
);
export const enginesReverse = new Map(
  Array.from(mapping.engines, ([k, v]) => [v, k]),
);
export const versionsReverse = new Map(
  Array.from(mapping.versions, ([k, v]) => [v, k]),
);
