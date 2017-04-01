int reduce(int i)  
{
  int x;
  switch( i ) 
  {
      case -1:
          x++;
          break;
      case 0 :
          x+=50;
          break;
      case 1 :
          x+=-100;
          break;
  }
  return x;

}
