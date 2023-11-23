/* eslint-disable */
export default {
  displayName: 'typescript-nest-shared-interceptors-wrapper-data',
  preset: '../../../../../../jest.preset.js',
  testEnvironment: 'node',
  transform: {
    '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
  },
  moduleFileExtensions: ['ts', 'js', 'html'],
  coverageDirectory:
    '../../../../../../coverage/libs/typescript/nest/shared/interceptors/wrapper-data',
};
