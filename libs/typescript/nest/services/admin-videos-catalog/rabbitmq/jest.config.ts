/* eslint-disable */
export default {
  displayName: 'typescript-nest-services-admin-videos-catalog-rabbitmq',
  preset: '../../../../../../jest.preset.js',
  testEnvironment: 'node',
  transform: {
    '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
  },
  moduleFileExtensions: ['ts', 'js', 'html'],
  coverageDirectory:
    '../../../../../../coverage/libs/typescript/nest/services/admin-videos-catalog/rabbitmq',
};
