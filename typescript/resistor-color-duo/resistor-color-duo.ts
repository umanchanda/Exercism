export class ResistorColor {
  private colorValues: string[] = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']
  private colors: string[];

  constructor(colors: string[]) {
    this.colors = colors.slice(0,2);

    if (this.colors.length < 2) {
      throw new Error("At least two colors need to be present")
    }
  }
  value = (): number => {
    let resistance: string = "";

    let colorOne: string = this.colorValues.indexOf(this.colors[0]).toString()
    let colorTwo: string = this.colorValues.indexOf(this.colors[1]).toString()

    resistance = colorOne.concat(colorTwo)

    return Number(resistance)
  };

  
}
