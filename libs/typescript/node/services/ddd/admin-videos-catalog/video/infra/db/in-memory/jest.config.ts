/* eslint-disable */
export default {
  displayName:
    'typescript-node-services-ddd-admin-videos-catalog-video-infra-db-in-memory',
  preset: '../../../../../../../../../../jest.preset.js',
  testEnvironment: 'node',
  transform: {
    '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
  },
  moduleFileExtensions: ['ts', 'js', 'html'],
  coverageDirectory:
    '../../../../../../../../../../coverage/libs/typescript/node/services/ddd/admin-videos-catalog/video/infra/db/in-memory',
};
