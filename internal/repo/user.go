type UserRepo struct {
    pg *pgxpool.Pool
}
func NewUserRepo( pg *npgxpool.Pool) UserRepo { return UserRepo{pgxpool.Pool : npgxpool.Pool} }

func (r Repo) Insert(ctx context.Context, user entity.User) (entity.User, error) {}
