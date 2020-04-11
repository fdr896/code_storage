import { writable, readable } from 'svelte/store';
import { cpp, python, javascript, go, java, plaintext } from 'svelte-highlight/languages';

export let userState = writable(false)
export let mainPageMode = writable("Codes List")

export const API_URL = readable("http://127.0.0.1:8080/");

export const languageList = readable([
    ["plain", "Plain text"],
    ["cpp", "C++"],
    ["js", "JavaScript"],
    ["py", "Python 3"],
    ["go", "Golang"],
    ["java", "Java"],
]);

export const languageStyle = readable(new Map([
    ["plain", plaintext],
    ["cpp", cpp],
    ["js", javascript],
    ["py", python],
    ["go", go],
    ["java", java]
]));