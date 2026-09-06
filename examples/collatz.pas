program Collatz;

var
  n: integer;
  steps: integer;
  largest: integer;

begin
  n := 27;
  steps := 0;
  largest := n;

  writeln('Collatz sequence starting at ', n, ':');
  writeln(n);

  while n <> 1 do
  begin
    if n mod 2 = 0 then
      n := n div 2
    else
      n := 3 * n + 1;

    if n > largest then
      largest := n;

    steps := steps + 1;
    writeln(n);
  end;

  writeln('Steps: ', steps);
  writeln('Largest value: ', largest);
end.
