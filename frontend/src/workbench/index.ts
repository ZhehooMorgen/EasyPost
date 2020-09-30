import Base from './Base'

export { Base }
export interface ITabOpener<T> {
    (data: T, Renderer: { new(): Base<T> }): void
}


class Animal<T> {
    makeSound(): void{}
}
class Dog extends Animal<number> {
    makeSound(): void {
        console.log('woof')
    }
}
class Cat extends Animal<string> {
    makeSound(): void {
        console.log('meow')
    }
}
const getAnimal = (name: string): { new(): Animal<number> } => {
    if (name === 'cat') return Cat
    return Dog
}

const getAnimalInstance = (name: string): Animal<number> => {
    if (name === 'cat') return new Cat()
    return new Dog()
}

console.log('sound:', new (getAnimal('dog'))().makeSound())
console.log('sound:', getAnimalInstance('cat').makeSound())