/**
 * Returns the error string from an error.
 *
 * @param error - Error object
 *
 * @returns The error string obtained from the error object.
 */
export function GetErrorMessage(error: unknown): string {
  if (error instanceof Error) return error.message;
  return String(error);
}
