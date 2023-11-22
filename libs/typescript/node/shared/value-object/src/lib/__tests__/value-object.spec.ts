import { ValueObject } from '../value-object';

class StringValueObject extends ValueObject {
  constructor(readonly value: string) {
    super();
  }
}

class ComplexValueObject extends ValueObject {
  constructor(readonly prop1: string, readonly prop2: number) {
    super();
  }
}

describe('ValueObject Unit Tests', () => {
  test('Should compare two value objects', () => {
    const vo1 = new StringValueObject('value');
    const vo2 = new StringValueObject('value');
    expect(vo1.equals(vo2)).toBe(true);

    const vo3 = new ComplexValueObject('value', 1);
    const vo4 = new ComplexValueObject('value', 1);
    expect(vo3.equals(vo4)).toBe(true);
  });

  test('Should compare two value objects with different values', () => {
    const vo1 = new StringValueObject('value');
    const vo2 = new StringValueObject('value2');
    expect(vo1.equals(vo2)).toBe(false);

    const vo3 = new ComplexValueObject('value', 1);
    const vo4 = new ComplexValueObject('value', 2);
    expect(vo3.equals(vo4)).toBe(false);
  });
});
