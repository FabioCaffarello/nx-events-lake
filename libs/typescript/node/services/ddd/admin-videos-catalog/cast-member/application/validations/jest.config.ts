/* eslint-disable */
export default {
  displayName:
    'typescript-node-services-ddd-admin-videos-catalog-cast-member-application-validations',
  preset: '../../../../../../../../../jest.preset.js',
  testEnvironment: 'node',
  transform: {
    '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
  },
  moduleFileExtensions: ['ts', 'js', 'html'],
  coverageDirectory:
    '../../../../../../../../../coverage/libs/typescript/node/services/ddd/admin-videos-catalog/cast-member/application/validations',
    // setupFilesAfterEnv: ['../../../../../../../../../expect-helpers.ts']
};
