/** concatArrays used for combining two transactions
 *
 * @param a - first uint8 array
 * @param b - second uint8 array
 * @returns returns a concatenated uint8 array
 */
 export const concatArrays = (a: any, b: any) => {
    let c = new Uint8Array(a.length + b.length);
    c.set(a);
    c.set(b, a.length);
    return c;
  };