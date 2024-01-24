/* eslint-disable */
export default {
  displayName: 'typescript-node-shared-ddd-utils-infra-message-broker',
  preset: '../../../../../../../jest.preset.js',
  testEnvironment: 'node',
  transform: {
    '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
  },
  moduleFileExtensions: ['ts', 'js', 'html'],
  coverageDirectory:
    '../../../../../../../coverage/libs/typescript/node/shared/ddd-utils/infra/message-broker',
};
