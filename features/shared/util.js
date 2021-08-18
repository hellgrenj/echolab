export const elementExist = (elementId) => {
    const el = document.getElementById(elementId)
    return el !== undefined && el !== null;
}