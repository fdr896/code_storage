export function AddPage(pageMode) {
    return pageMode === 'Add Code';
}

export function ViewPage(pageMode) {
    return pageMode === 'Codes List' ||
           pageMode === 'Show Code'  ||
           pageMode === 'Edit Code';
}