program MultiplicationGrid;

var
  i, j, value: integer;

begin
  writeln('Multiplication Grid');
  writeln('------------------');

  for i := 1 to 10 do
  begin
    for j := 1 to 10 do
    begin
      value := i * j;
      write(value, ' ');
    end;
    writeln;
  end;

  writeln;
  writeln('Done!');
end.

