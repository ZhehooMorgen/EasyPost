import Base from './Base'

export { Base }
export interface ITabOpener<T> {
    (data: T, Renderer: Base<T>): void
}

// interface Processor<T> {
//     (data: T): void
// }

// let textProcessor = (data: string) => { }

// function Process<T>(data: T, processor: Processor<T>) {
//     processor(data)
// }

// Process("sdfs", textProcessor)