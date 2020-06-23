package game

import (
	"testing"
)

func test_that_team_is_set_up_correctly(
	t *testing.T,
	color string, team *Team,
	q, p, h, c, e, g, k Piece,
) {
	if team.Q != q {
		t.Errorf("%sTeam.Q is not set up correctly\n", color)
	}
	if team.P != p {
		t.Errorf("%sTeam.P is not set up correctly\n", color)
	}
	if team.H != h {
		t.Errorf("%sTeam.H is not set up correctly\n", color)
	}
	if team.C != c {
		t.Errorf("%sTeam.C is not set up correctly\n", color)
	}
	if team.E != e {
		t.Errorf("%sTeam.E is not set up correctly\n", color)
	}
	if team.G != g {
		t.Errorf("%sTeam.G is not set up correctly\n", color)
	}
	if team.K != k {
		t.Errorf("%sTeam.K is not set up correctly\n", color)
	}
	if team.QPHCEGK[0] != q {
		t.Errorf("%sTeam.QPHCEGK[0] is not set up correctly\n", color)
	}
	if team.QPHCEGK[1] != p {
		t.Errorf("%sTeam.QPHCEGK[1] is not set up correctly\n", color)
	}
	if team.QPHCEGK[2] != h {
		t.Errorf("%sTeam.QPHCEGK[2] is not set up correctly\n", color)
	}
	if team.QPHCEGK[3] != c {
		t.Errorf("%sTeam.QPHCEGK[3] is not set up correctly\n", color)
	}
	if team.QPHCEGK[4] != e {
		t.Errorf("%sTeam.QPHCEGK[4] is not set up correctly\n", color)
	}
	if team.QPHCEGK[5] != g {
		t.Errorf("%sTeam.QPHCEGK[5] is not set up correctly\n", color)
	}
	if team.QPHCEGK[6] != k {
		t.Errorf("%sTeam.QPHCEGK[6] is not set up correctly\n", color)
	}
}

func TestTeams(t *testing.T) {
	test_that_team_is_set_up_correctly(t, "Red", &RedTeam,
		RedCannon,
		RedPawn,
		RedHorse, RedCart, RedElephant, RedGuard,
		RedKing,
	)
	test_that_team_is_set_up_correctly(t, "Black", &BlackTeam,
		BlackCannon,
		BlackPawn,
		BlackHorse, BlackCart, BlackElephant, BlackGuard,
		BlackKing,
	)
}
